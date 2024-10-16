FROM rust:1.81

WORKDIR /app
ENV CARGO_BUILD_TARGET_DIR=/tmp/target
RUN rustup override set nightly 

CMD ["cargo", "run"]