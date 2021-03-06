package storage

// Code generated by chromedp-gen. DO NOT EDIT.

import (
	cdp "github.com/knq/chromedp/cdp"
)

// EventCacheStorageListUpdated a cache has been added/deleted.
type EventCacheStorageListUpdated struct {
	Origin string `json:"origin"` // Origin to update.
}

// EventCacheStorageContentUpdated a cache's contents have been modified.
type EventCacheStorageContentUpdated struct {
	Origin    string `json:"origin"`    // Origin to update.
	CacheName string `json:"cacheName"` // Name of cache in origin.
}

// EventIndexedDBListUpdated the origin's IndexedDB database list has been
// modified.
type EventIndexedDBListUpdated struct {
	Origin string `json:"origin"` // Origin to update.
}

// EventIndexedDBContentUpdated the origin's IndexedDB object store has been
// modified.
type EventIndexedDBContentUpdated struct {
	Origin          string `json:"origin"`          // Origin to update.
	DatabaseName    string `json:"databaseName"`    // Database to update.
	ObjectStoreName string `json:"objectStoreName"` // ObjectStore to update.
}

// EventTypes all event types in the domain.
var EventTypes = []cdp.MethodType{
	cdp.EventStorageCacheStorageListUpdated,
	cdp.EventStorageCacheStorageContentUpdated,
	cdp.EventStorageIndexedDBListUpdated,
	cdp.EventStorageIndexedDBContentUpdated,
}
