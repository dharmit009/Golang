# Hash Table:

Hash tables are *key-value* pairs. Each key has to be unique
in hash table.

* Each value is associated with a unique key.
* `Hash Function` is used to compute the slot of a key.
* Indexing is done using [].

*🚩 Note:* We never call this Hash function explicitly.
All the underlying things are carried out by Go itself.

**Here is an example of the same:**

| Key  | Value |
|------|-------|
| Joe  | x     |
| Jane | y     |
| Pat  | z     |

**Advantages:**

* Faster Lookup than lists.
    ** Constant-time vs Linear-time.
* Arbitrary Keys.
    ** Not Ints, like slices or arrays.

**Disadvantages:**

* You may have collisions
    ** Two keys hash to same slot. ( ⚠️ Collisions are rare)
