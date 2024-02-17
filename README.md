
# LRU Cache Implementation in Go

This is an implementation of a Least Recently Used (LRU) cache in Go.

## Overview

The LRU cache is a data structure used for caching where the least recently used items are removed when the cache reaches its capacity. This implementation consists of the following files:

- **main.go**: Contains the main function and user interface for interacting with the LRU cache.
- **LRUcache.go**: Defines the `Cache` struct and its methods for setting, getting, printing the state, and clearing the cache.
- **operation.go**: Implements operations related to cache manipulation, such as adding, deleting, and moving nodes within the cache.


## Functionality

- `Set`: Adds a key-value pair to the cache. If the key already exists, its value is updated. If the cache is full, the least recently used item is removed.
- `Get`: Retrieves the value associated with the given key from the cache. If the key exists, its corresponding value is returned, and the item is moved to the front of the cache. If the key does not exist, an appropriate message is displayed.
- `PrintState`: Prints the current state of the cache, including all key-value pairs.
- `ClearCache`: Clears all items from the cache.
