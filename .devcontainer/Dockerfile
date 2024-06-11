FROM mcr.microsoft.com/devcontainers/go:1-1.22-bookworm

# Prepare filesystem
WORKDIR /workspaces/daneel
RUN mkdir -p /home/vscode/wordlists && \
    mkdir -p /home/vscode/.config/subfinder

# Install Go tools
RUN go install github.com/spf13/cobra-cli@latest && \
    go install github.com/owasp-amass/amass/v4/...@master && \
    go install github.com/projectdiscovery/subfinder/v2/cmd/subfinder@latest

# seclists
RUN cd /home/vscode/wordlists \
    && git clone --depth 1 https://github.com/danielmiessler/SecLists.git

# Set up config
COPY config/subfinder-provider-config.yaml /home/vscode/.config/subfinder/provider-config.yaml