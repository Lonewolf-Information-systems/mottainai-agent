/*

Copyright (C) 2017-2018  Ettore Di Giacinto <mudler@gentoo.org>
Credits goes also to Gogs authors, some code portions and re-implemented design
are also coming from the Gogs project, which is using the go-macaron framework
and was really source of ispiration. Kudos to them!

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.

*/

package database

import "github.com/MottainaiCI/mottainai-server/pkg/nodes"

var NodeColl = "Nodes"

func (d *Database) CreateNode(t map[string]interface{}) (int, error) {
	return d.InsertDoc(NodeColl, t)
}

func (d *Database) DeleteNode(docID int) error {
	return d.DeleteDoc(NodeColl, docID)
}

func (d *Database) UpdateNode(docID int, t map[string]interface{}) error {
	return d.UpdateDoc(NodeColl, docID, t)
}

func (d *Database) GetNode(docID int) (nodes.Node, error) {
	doc, err := d.GetDoc(NodeColl, docID)
	if err != nil {
		return nodes.Node{}, err
	}
	t := nodes.NewNodeFromMap(doc)
	t.ID = docID
	return t, err
}

func (d *Database) ListNodes() []DocItem {
	return d.ListDocs(NodeColl)
}

func (d *Database) AllNodes() []nodes.Node {
	nodec := d.DB().Use(NodeColl)
	node_list := make([]nodes.Node, 0)

	nodec.ForEachDoc(func(id int, docContent []byte) (willMoveOn bool) {
		t := nodes.NewFromJson(docContent)
		t.ID = id
		node_list = append(node_list, t)
		return true
	})
	return node_list
}
