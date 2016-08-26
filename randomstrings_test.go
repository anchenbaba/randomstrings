/*
* @Author: anchen
* @Date:   2016-08-26 13:34:56
* @Last Modified by:   anchen
* @Last Modified time: 2016-08-26 13:56:24
 */

package randomstrings

import (
	"testing"
)

func TestRandomString(t *testing.T) {
	var r []byte

	r = RandomString(10, 1)
	t.Log("Sucess: ", string(r))

}
