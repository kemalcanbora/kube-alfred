# Kube-Alfred

Kube-Alfred is an AI-powered assistant for Kubernetes. It allows you to ask questions about Kubernetes directly from kubectl, providing intelligent responses using the Anthropic API.

## Installation

### Using Krew

If you have [Krew](https://krew.sigs.k8s.io/) installed, you can easily install kube-alfred:

```bash
kubectl krew install alfred
```

### Manual Installation

If you prefer to install manually:

1. Download the appropriate version for your operating system and architecture from the [releases page](https://github.com/kemalcanbora/kube-alfred/releases).
2. Rename the downloaded file to `kube-alfred`.
3. Make it executable: `chmod +x kube-alfred`
4. Move it to a directory in your PATH, e.g., `mv kube-alfred /usr/local/bin/`

## Usage

Once installed, you can use kube-alfred by running:

```bash
kubectl alfred "<your question about Kubernetes>"
```

For example:

```bash
kubectl alfred "How do I create a deployment?"
```

## Configuration

Before using kube-alfred, you need to set up your Anthropic API key. You can do this by setting an environment variable:

```bash
export ANTHROPIC_API_KEY=your_api_key_here
```

For persistent configuration, add this line to your shell's configuration file (e.g., `~/.bashrc` or `~/.zshrc`).

## Examples

Here are a few example questions you can ask:

```bash
kubectl alfred "What is a Pod in Kubernetes?"
kubectl alfred "How do I troubleshoot a failing deployment?"
kubectl alfred "Explain Kubernetes Services and their types"
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

[MIT License](LICENSE)

## Support

If you encounter any issues or have questions, please file an issue on the [GitHub repository](https://github.com/kemalcanbora/kube-alfred/issues).