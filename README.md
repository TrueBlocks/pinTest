# IPFS Pin Test

1. Make a package that we can use with these endpoings:
   - **manifest.Add(chain, range)** pins both Bloom and Index (does not zip) and returns struct { Range:, BloomHash, IndexHash }
   - **manifest.PinAll(chain)** scans blooms and index for chain, pinning everything
   - **manifest.ListPinata(all|chain)**
