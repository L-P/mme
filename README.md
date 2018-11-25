# mme - Majora's Mask Explorer

## Usage
```
./mme ROM
# Open http://localhost:8064 in your favorite web browser.
```

## Development
A file named `rom.z64` is required at the repository root.

1. Run the front dev server: `cd front && yarn serve`
2. Run the back dev server: `make run`
3. Open your browser to http://localhost:8080

## Production
1. `make`
2. `./mme ROM`
3. Open your browser to http://localhost:8064

URIs and ports are hardcoded for now.
