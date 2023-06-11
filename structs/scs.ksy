meta:
  id: scs
  file-extension: scs
  endian: le
seq:
  - id: magic
    size: 4
    contents: SCS#
  - id: version
    type: u2
  - id: salt
    type: u2
  - id: hash_method
    size: 4
    contents: CITY
  - id: entry_count
    type: u4
  - id: start_offset
    type: u4
  - id: padding
    size: 12
  - id: entries
    type: entry
    repeat: expr
    repeat-expr: entry_count
types:
  entry:
    seq:
      - id: hash
        type: u8
      - id: offset
        type: u8
      - id: flags
        type: u4
      - id: crc
        type: u4
      - id: size
        type: u4
      - id: compressed_size
        type: u4
    instances:
      is_dir:
        value: (flags & (1 << 0)) != 0
      is_compressed:
        value: (flags & (1 << 1)) != 0
      is_verify:   
        value: (flags & (1 << 2)) != 0
      is_encrypted:
        value: (flags & (1 << 3)) != 0
      len:
        value: > 
          is_compressed ? compressed_size : size
      dir:
        io: _root._io
        pos: offset
        size: len
        type: str
        encoding: ASCII
        if: is_dir
      file:
        io: _root._io
        pos: offset
        size: len
        if: not is_dir
        
          
        
  