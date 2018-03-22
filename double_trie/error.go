/************************************************************
** @Description: double_trie
** @Author: george hao
** @Date:   2018-03-21 15:16
** @Last Modified by:  george hao
** @Last Modified time: 2018-03-21 15:16
*************************************************************/
package double_trie

import "errors"

var (
	ErrInvalidDataType = errors.New("dt: invalid datatype")
	ErrInvalidValue    = errors.New("dt: invalid value")
	ErrInvalidKey      = errors.New("dt: invalid key")
	ErrNoPath          = errors.New("dt: no path")
	ErrNoValue         = errors.New("dt: no value")
)
