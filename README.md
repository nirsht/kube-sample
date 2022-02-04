# kube-sample

### Requirments

- kubectl

### Dev requirments

- golang version > 1.17

### Installation

```bash
    go install github.com/nirsht/kube-sample
```

### Installation in development

```bash
    go mod download
    go build
    go install
```

### Add new resource

There are 3 steps in order to add new resource:

1. Add YAML file in manifest directory
2. Run bash file
   ```bash
    chmod +x kubernetes_supported_resources/manifest/YamlToGO.sh
    kubernetes_supported_resources/manifest/YamlToGO.sh
   ```
3. In kubernetes_supported_resources/types there is 'resourcesMap' var, add there the new resource to the dictionary
