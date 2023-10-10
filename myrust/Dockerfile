# Build Stage
FROM rust:latest as builder

WORKDIR /app

# Copying and dummy build to cache dependencies
COPY Cargo.toml Cargo.lock ./

# Copy source code and build again
COPY ./src ./src
RUN cargo build --release

# Runtime Stage
FROM rust:slim

# You might not need to install libssl-dev if the base image already has it.
# But if you have other dependencies, install them here.

WORKDIR /app

# Copying binary from builder stage
COPY --from=builder /app/target/release/myrust .

CMD ["./myrust"]
