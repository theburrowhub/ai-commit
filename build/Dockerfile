FROM alpine:3.20.3

LABEL org.opencontainers.image.source=https://github.com/theburrowhub/ai-commit
LABEL org.opencontainers.image.description="Commit meesages using AI"
LABEL org.opencontainers.image.licenses=MIT

RUN apk --no-cache add \
    ca-certificates \
    openssh-client \
    git \
    vim \
    bash

RUN addgroup -S ai-commit && adduser -S ai-commit -G ai-commit

COPY ai-commit /usr/local/bin/ai-commit

# Mount the source code here
WORKDIR /source

USER ai-commit

RUN echo "git config --global --add safe.directory /source" >> /home/ai-commit/.bashrc && \
    echo 'git config --global user.name "$GIT_USER_NAME"' >> /home/ai-commit/.bashrc && \
    echo 'git config --global user.email "$GIT_USER_EMAIL"' >> /home/ai-commit/.bashrc

# Set the entrypoint to start the SSH agent
ENTRYPOINT ["/bin/bash", "-c", "source /home/ai-commit/.bashrc && /usr/local/bin/ai-commit $@", "--"]
CMD [""]