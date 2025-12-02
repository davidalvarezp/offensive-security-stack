masscan_fast_scan.sh
#!/bin/bash
# masscan_fast_scan.sh - Fast port scan with Masscan
TARGET=\$1
OUTPUT_DIR="./\${TARGET}"
mkdir -p \$OUTPUT_DIR
masscan \$TARGET -p1-65535 --rate=1000 -oJ \$OUTPUT_DIR/masscan_\${TARGET}.json
