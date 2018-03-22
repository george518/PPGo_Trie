/************************************************************
** @Description: double_trie
** @Author: george hao
** @Date:   2018-03-21 15:14
** @Last Modified by:  george hao
** @Last Modified time: 2018-03-21 15:14
*************************************************************/
package double_trie

func (da *Dat) Insert(key []byte, value int) error {
	if value < 0 || value >= ValueLimit {
		return ErrInvalidValue
	}
	p := da.get(key, 0, 0)
	*p = value
	return nil
}

func (da *Dat) Jump(path []byte, from int) (to int, err error) {
	for _, b := range path {
		if da.Array[from].Value >= 0 {
			return from, ErrNoPath
		}
		to = da.Array[from].base() ^ int(b)
		if da.Array[to].Check != from {
			return from, ErrNoPath
		}
		from = to
	}
	return to, nil
}

func (da *Dat) Key(id int) (key []byte, err error) {
	for id > 0 {
		from := da.Array[id].Check
		if from < 0 {
			return nil, ErrNoPath
		}
		if char := byte(da.Array[from].base() ^ id); char != 0 {
			key = append(key, char)
		}
		id = from
	}
	if id != 0 || len(key) == 0 {
		return nil, ErrInvalidKey
	}
	for i := 0; i < len(key)/2; i++ {
		key[i], key[len(key)-i-1] = key[len(key)-i-1], key[i]
	}
	return key, nil
}

func (da *Dat) Value(id int) (value int, err error) {
	value = da.Array[id].Value
	if value >= 0 {
		return value, nil
	}
	to := da.Array[id].base()
	if da.Array[to].Check == id && da.Array[to].Value >= 0 {
		return da.Array[to].Value, nil
	}
	return 0, ErrNoValue
}

func (da *Dat) Update(key []byte, value int) error {
	p := da.get(key, 0, 0)

	// key was not inserted
	if *p == ValueLimit {
		*p = value
		return nil
	}

	// key was inserted before
	if *p+value < 0 || *p+value >= ValueLimit {
		return ErrInvalidValue
	}
	*p += value
	return nil
}

func (da *Dat) Delete(key []byte) error {
	// if the path does not exist, or the end is not a leaf, nothing to delete
	to, err := da.Jump(key, 0)
	if err != nil {
		return ErrNoPath
	}

	if da.Array[to].Value < 0 {
		base := da.Array[to].base()
		if da.Array[base].Check == to {
			to = base
		}
	}

	for to > 0 {
		from := da.Array[to].Check
		base := da.Array[from].base()
		label := byte(to ^ base)

		// if `to` has sibling, remove `to` from the sibling list, then stop
		if da.Ninfos[to].Sibling != 0 || da.Ninfos[from].Child != label {
			// delete the label from the child ring first
			da.popSibling(from, base, label)
			// then release the current node `to` to the empty node ring
			da.pushEnode(to)
			break
		}
		// otherwise, just release the current node `to` to the empty node ring
		da.pushEnode(to)
		// then check its parent node
		to = from
	}
	return nil
}

// Get returns the value associated with the given `key`.
// It is equivalent to
//		id, err1 = Jump(key)
//		value, err2 = Value(id)
// Thus, it may return ErrNoPath or ErrNoValue,
func (da *Dat) Get(key []byte) (value int, err error) {
	to, err := da.Jump(key, 0)
	if err != nil {
		return 0, err
	}
	return da.Value(to)
}
