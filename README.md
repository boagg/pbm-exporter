# pbm-exporter

Prometheus exporter for PBM (Percona Backup MongoDB) written in Go.

## 🔄 Fork et migration vers Go

Ce projet est un **fork migré vers Go** du [pbm-exporter original en Node.js](https://github.com/koumoul-dev/pbm-exporter). 

### 🎯 Pourquoi cette migration ?

La migration vers Go suit les **standards de l'écosystème Prometheus** :

- **📈 Cohérence** : Prometheus et la quasi-totalité des exporteurs officiels sont écrits en Go
- **🚀 Performance** : Binaire statique léger (~15MB) vs runtime Node.js (~100MB+)
- **📦 Distribution simplifiée** : Un seul fichier exécutable, sans dépendances
- **⚡ Démarrage instantané** : Plus de temps d'initialisation du runtime JavaScript
- **🏗️ Build natif** : Cross-compilation native pour toutes les plateformes

Cette approche garantit une **meilleure intégration** dans l'écosystème de monitoring moderne et suit les recommandations de l'équipe Prometheus.

> **Repository original** : [percona/pbm-exporter](https://github.com/koumoul-dev/pbm-exporter) (Node.js)

## 🚀 Installation rapide

### Via GitHub Releases (Recommandé)

**Installation automatique (Linux/macOS):**
```bash
curl -fsSL https://github.com/boagg/pbm-exporter/releases/latest/download/install.sh | bash
```

**Installation manuelle:**

1. **Téléchargez le binaire** depuis [GitHub Releases](https://github.com/boagg/pbm-exporter/releases)
2. **Linux/macOS:**
   ```bash
   # Remplacez par votre plateforme (linux-amd64, linux-arm64, darwin-amd64, darwin-arm64)
   curl -fsSL https://github.com/boagg/pbm-exporter/releases/latest/download/pbm-exporter-linux-amd64.tar.gz | tar -xz
   sudo mv pbm-exporter-linux-amd64/pbm-exporter /usr/local/bin/
   pbm-exporter --version
   ```

3. **Windows:**
   ```powershell
   # Téléchargez et extrayez pbm-exporter-windows-amd64.zip
   Invoke-WebRequest -Uri "https://github.com/boagg/pbm-exporter/releases/latest/download/pbm-exporter-windows-amd64.zip" -OutFile "pbm-exporter.zip"
   Expand-Archive -Path "pbm-exporter.zip" -DestinationPath "."
   ```

### Via Docker

```bash
docker run -p 9216:9216 ghcr.io/boagg/pbm-exporter:latest
```

## Features

- ✅ **Migré vers Go** : Fork du projet original pour suivre les standards Prometheus
- ✅ **Zero dependencies**: Single binary with no external dependencies
- ✅ **Cross-platform**: Available for Linux, macOS, and Windows (amd64/arm64)
- ✅ **Docker support**: Multi-stage build with minimal Alpine image
- ✅ **Systemd integration**: Easy installation with systemd service
- ✅ **Prometheus metrics**: Full compatibility with Prometheus monitoring
- ✅ **Health checks**: Built-in health check endpoint
- ✅ **Graceful shutdown**: Proper signal handling
- ✅ **Lightweight**: ~15MB binary vs ~100MB+ with Node.js runtime
- ✅ **Fast startup**: Instant startup time

## Metrics

The exporter provides the following Prometheus metrics:

- `pbm_snapshots_total{status}` - Number of snapshots per status
- `pbm_snapshots{name,status}` - Detail of snapshots with statuses  
- `pbm_last_snapshot{status}` - Status of last snapshot
- `pbm_last_snapshot_error` - 1 if last snapshot is in error
- `pbm_last_snapshot_since_seconds` - Time since last snapshot
- `pbm_nodes_total{status}` - Number of nodes per status
- `pbm_nodes{rs,host,status}` - Detail of nodes with statuses
- `pbm_pitr_chunks_total` - Number of PITR chunks
- `pbm_pitr_error` - 1 if PITR is in error
- `pbm_last_pitr_chunk_since_seconds` - Time since last PITR chunk

## Quick Start

### 1. Installation

Choisissez votre méthode d'installation préférée :

- **GitHub Releases** (voir section ci-dessus) - Recommandé
- **Docker** : `docker run -p 9216:9216 ghcr.io/boagg/pbm-exporter:latest`
- **Build from source** : voir section "Development" ci-dessous

### 2. Configuration

**Option A: Variables d'environnement**
```bash
export PBM_MONGODB_URI="mongodb://localhost:27017"
export PBM_LISTEN_PORT="9216"
pbm-exporter
```

**Option B: Fichier de configuration**
```bash
# Créer /etc/default/pbm-exporter ou ./config.env
echo 'PBM_MONGODB_URI=mongodb://localhost:27017' > config.env
pbm-exporter
```

### 3. Vérification

```bash
# Test du endpoint de métriques
curl http://localhost:9216/metrics

# Test du health check
curl http://localhost:9216/health
```

## Installation avancée

### Download and Install from Source

1. **Build from source:**
```bash
git clone https://github.com/boagg/pbm-exporter.git
cd pbm-exporter
make build
```

2. **Install with systemd service:**
```bash
sudo ./install.sh --start
```

3. **Configure MongoDB URI:**
```bash
sudo vi /etc/default/pbm-exporter
# Set: PBM_MONGODB_URI=mongodb://your-mongodb:27017
sudo systemctl restart pbm-exporter
```

### Docker

```bash
# Build and run with docker-compose
docker-compose up -d

# Or run manually
docker build -t pbm-exporter .
docker run -d \
  -p 9090:9090 \
  -e PBM_MONGODB_URI=mongodb://mongodb:27017 \
  pbm-exporter
```

### Binary Usage

```bash
# Set MongoDB URI
export PBM_MONGODB_URI=mongodb://localhost:27017

# Run the exporter
./pbm-exporter

# Or specify port
PORT=8080 ./pbm-exporter
```

## Configuration

The exporter is configured via environment variables:

- `PBM_MONGODB_URI` (required) - MongoDB connection URI
- `PORT` (optional) - Port to listen on (default: 9090)

## Build Options

```bash
# Build for current platform
make build

# Build with debug info
make build-debug

# Cross-compile for all platforms
make cross-compile

# Create release archives
make release

# Build Docker image
make docker

# Install to system
make install

# Development mode with auto-rebuild
make dev PBM_MONGODB_URI=mongodb://localhost:27017
```

## Installation Options

### 1. System Installation (Recommended)

```bash
# Build and install
make build
sudo ./install.sh --start

# Configure
sudo vi /etc/default/pbm-exporter
sudo systemctl restart pbm-exporter

# Monitor
sudo systemctl status pbm-exporter
sudo journalctl -u pbm-exporter -f
```

### 2. Manual Installation

```bash
# Build binary
make build

# Copy to system
sudo cp build/pbm-exporter /usr/local/bin/
sudo chmod +x /usr/local/bin/pbm-exporter

# Create systemd service (see install.sh for example)
```

### 3. Docker Installation

```bash
# Using docker-compose
docker-compose up -d pbm-exporter

# Using Docker directly
docker run -d \
  --name pbm-exporter \
  -p 9090:9090 \
  -e PBM_MONGODB_URI=mongodb://mongodb:27017 \
  pbm-exporter:latest
```

## Development

### Prerequisites

- Go 1.21 or later
- MongoDB with PBM configured
- Make (optional, for convenience)

### Setup Development Environment

```bash
# Clone repository
git clone https://github.com/your-org/pbm-exporter.git
cd pbm-exporter

# Install dependencies
go mod download

# Run tests
make test

# Development mode (auto-rebuild on changes)
make dev PBM_MONGODB_URI=mongodb://localhost:27017
```

## Testing with PBM

Run PBM agent and MongoDB containers:

```bash
# Create network
docker network create pbm-exporter-test

# Start services
docker-compose up -d

# Initialize replica set
docker-compose exec mongo mongo
>> rs.initiate({_id: 'pbm-exporter-test', members: [{_id: 0, host: 'mongo:27017'}]})
>> db.test.insert({'test': 'Test !!'})

# Access metrics
curl http://localhost:9090/metrics
```

## Monitoring Integration

### Prometheus Configuration

```yaml
# prometheus.yml
scrape_configs:
  - job_name: 'pbm-exporter'
    static_configs:
      - targets: ['localhost:9090']
    scrape_interval: 30s
    metrics_path: /metrics
```

### Grafana Dashboard

Import the provided Grafana dashboard (dashboard.json) or create custom panels using the available metrics.

## Troubleshooting

### Service Not Starting

```bash
# Check service status
sudo systemctl status pbm-exporter

# View logs
sudo journalctl -u pbm-exporter -f

# Check configuration
cat /etc/default/pbm-exporter
```

### Connection Issues

```bash
# Test MongoDB connection
mongo $PBM_MONGODB_URI

# Check firewall
sudo netstat -tlnp | grep 9090

# Test exporter manually
PBM_MONGODB_URI=mongodb://localhost:27017 ./pbm-exporter
```

### Metrics Issues

```bash
# Test metrics endpoint
curl http://localhost:9090/metrics

# Check PBM collections in MongoDB
mongo $PBM_MONGODB_URI
>> use admin
>> show collections
>> db.pbmBackups.count()
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Run `make test`
6. Submit a pull request

## License

Apache License 2.0 - see LICENSE file for details.

## Changelog

### v2.0.0+ (Go Fork)
- ✅ Fork du [projet original](https://github.com/percona/pbm-exporter) et réécriture complète en Go
- ✅ Single binary with no dependencies
- ✅ Cross-platform support (Linux, macOS, Windows, ARM64)
- ✅ Improved performance and memory usage
- ✅ Better error handling and logging
- ✅ Systemd integration
- ✅ Enhanced Docker image with security hardening
- ✅ GitHub Actions CI/CD with automated releases

### Historique du projet original (Node.js)
- **v0.1.x** : Implémentation initiale en Node.js par Percona
- **Repository original** : [percona/pbm-exporter](https://github.com/percona/pbm-exporter)

Configure PBM and prepare first backup:

```
docker-compose exec pbm-agent bash
>> pbm config --file=/tmp/pbm-config.yaml
>> pbm backup
>> pbm config --set=pitr.enabled=true
```

Build and test the image:

```
docker build . -t pbm-exporter && docker run -it --rm -p 9090:9090 -e DEBUG=pbm-exporter -e PBM_MONGODB_URI=mongodb://mongo:27017 --network pbm-exporter-test --name pbm-exporter-test pbm-exporter
curl http://localhost:9090/metrics
```