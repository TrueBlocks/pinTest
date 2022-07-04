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


Questions to be answered:
--------------------------
1. Zipping the files means we can only create them on one operating system (either Linux or Mac) because of this (https://unix.stackexchange.com/questions/570477/gzip-doesnt-produce-the-same-compressed-result-on-macos-vs-linux). I have a feeling that this extends across other linuxes as well as across Mac/Linux. It may not be worth ZIPing. File size comparison:

| Type | Zipped | Not Zipped |
|----|----|----|
|Blooms|2.314 GB | 3.067 GB |
|Index|41.560 GB|74.238 GB|

2. What do domains do in Pinata - https://ipfs.unchainedindex.io/ipfs, https://sepolia.unchainedindex.io/ipfs. I have a feeling they don't do much.

3. Could we 'package up' larger files for the `chifra init --all` case? 0-1,000,000, 1,000,000-2,000,000, etc.
