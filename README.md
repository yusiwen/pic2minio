# pic2minio

Upload picture files to minio service, used as uploader for [Typora](https://typora.io/).

## Usage

### Configuration

Config location is: `$HOME/.config/pic2minio.yaml`

Sample configuration:

```yaml
endpoint: MINIO_SERVICE_HOST_NAME
access-key: ACCESS_KEY
secret-key: SECRET_KEY
bucket: BUCKET
base-dir: FOLDER_IN_BUCKET
```

### Standalone

```bash
G:\\apps\\pic2minio.exe pic-file-1 pic-file-2 ...
```

### In Typora

Setup [custom command in Typora](https://support.typora.io/Upload-Image/) as:

```bash
G:\\apps\\pic2minio.exe
```

