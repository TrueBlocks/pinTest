# IPFS Pin Test

1. Create a package called `pinning` or `ipfs`
   
2. Make a package that we can use with these endpoings:

| Function                   | Description                                                                                                          |
| -------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| **pinning.Add(fullPath)** | Given a full path to a file, pins the file locally (if IPFS is running) and remotely (assume user has the right key) |
