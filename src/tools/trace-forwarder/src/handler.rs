// Copyright (c) 2020-2021 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

use anyhow::{anyhow, Context, Result};
use byteorder::{ByteOrder, NetworkEndian};
use futures::executor::block_on;
use opentelemetry::sdk::export::trace::{SpanData, SpanExporter};
use slog::{debug, info, o, Logger};
use std::fs::File;
use std::io::{ErrorKind, Read};
use std::os::unix::io::{FromRawFd, RawFd};

// The VSOCK "packet" protocol used comprises two elements:
//
// 1) The header (the number of bytes in the payload).
// 2) The payload bytes.
//
// This constant defines the number of bytes used to encode the header on the
// wire. In other words, the first 64-bits of the packet contain a number
// specifying how many bytes are in the remainder of the packet.
//
// Must match the value of the variable of the same name in the agents
// vsock-exporter.
const HEADER_SIZE_BYTES: u64 = std::mem::size_of::<u64>() as u64;

fn mk_io_err(msg: &str) -> std::io::Error {
    std::io::Error::new(std::io::ErrorKind::Other, msg.to_string())
}

async fn handle_async_connection<'a>(
    logger: Logger,
    mut conn: &'a mut dyn Read,
    exporter: &'a mut dyn SpanExporter,
    dump_only: bool,
) -> Result<()> {
    let logger = logger.new(o!("subsystem" => "handler"));

    debug!(logger, "handling connection");

    handle_trace_data(logger.clone(), &mut conn, exporter, dump_only)
        .await
        .map_err(|e| mk_io_err(&format!("failed to handle data: {:}", e)))?;

    debug!(&logger, "handled connection");

    Ok(())
}

async fn handle_trace_data<'a>(
    logger: Logger,
    reader: &'a mut dyn Read,
    exporter: &'a mut dyn SpanExporter,
    dump_only: bool,
) -> Result<()> {
    loop {
        let mut header: [u8; HEADER_SIZE_BYTES as usize] = [0; HEADER_SIZE_BYTES as usize];

        info!(logger, "waiting for traces");

        match reader.read_exact(&mut header) {
            Ok(_) => debug!(logger, "read header"),
            Err(e) => {
                if e.kind() == ErrorKind::UnexpectedEof {
                    info!(logger, "agent shut down");
                    break;
                }

                return Err(anyhow!("failed to read header: {:}", e));
            }
        };

        let payload_len: u64 = NetworkEndian::read_u64(&header);

        let mut encoded_payload = Vec::with_capacity(payload_len as usize);
        encoded_payload.resize(payload_len as usize, 0);

        reader
            .read_exact(&mut encoded_payload)
            .with_context(|| "failed to read payload")?;

        debug!(logger, "read payload");

        let span_data: SpanData =
            bincode::deserialize(&encoded_payload[..]).expect("failed to deserialise payload");

        debug!(logger, "deserialised payload");

        if dump_only {
            debug!(logger, "dump-only: {:?}", span_data);
        } else {
            let batch = vec![span_data];

            // Call low-level Jaeger exporter to send the trace span immediately.
            let result = exporter.export(batch).await;

            if result.is_err() {
                return Err(anyhow!("failed to export trace spans: {:?}", result));
            }

            debug!(logger, "exported trace spans");
        }
    }

    Ok(())
}

pub fn handle_connection(
    logger: Logger,
    fd: RawFd,
    exporter: &mut dyn SpanExporter,
    dump_only: bool,
) -> Result<()> {
    let mut file = unsafe { File::from_raw_fd(fd) };

    let conn = handle_async_connection(logger, &mut file, exporter, dump_only);

    block_on(conn)?;

    Ok(())
}
