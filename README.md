# IPFS Pin Test

1. Create a package called `pinning` or `ipfs`
   
2. Make a package that we can use with these endpoings:

| Function                   | Description                                                                                                          |
| -------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| **pinning.Add(fullPath)** | Given a full path to a file, pins the file locally (if IPFS is running) and remotely (assume user has the right key) |
| **pinning.List(remoteName)** | List pins on remote service |
| **pinning.Remove(fullPath)** | Remove a pin from local or remote service |

These are totally off the top of my head. We can expand as needed.

------
An alternative to this would be to find a package that already exists and does these things. In that case, this repo can become an example of how to use such a package.
