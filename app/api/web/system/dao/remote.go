package dao

import (
	"encoding/json"
	"mall/app/api/web/system/model"

	"github.com/techleeone/gophp/serialize"
)

func (d *Dao) LoadRemote() model.Remote {
	var remote model.Remote
	remotePhpStr, err := d.settingLoad("remote")
	if err != nil {
		return remote
	}

	remoteMap, err := serialize.UnMarshal([]byte(remotePhpStr))
	if err != nil {
		return remote
	}

	data, err := json.Marshal(remoteMap)
	if err != nil {
		return remote
	}

	json.Unmarshal(data, &remote)

	return remote

}
