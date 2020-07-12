import 'dart:math';

import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:mobile/models/list/dish/dish-item.dart';

class DishListPage extends StatelessWidget {
  List<DishItem> _dishList;
  String _categoryName;
  String _categorySysName;

  DishListPage(String categoryName, String categorySysName) {
    this._categoryName = categoryName;
    this._categorySysName = categorySysName;
    this._dishList = this._loadDish(categorySysName);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(_categoryName),
      ),
      body: Container(
        child: ListView.builder(
          itemCount: _dishList != null ? _dishList.length : 0,
          itemBuilder: (context, index) {
            return Card(
              margin: EdgeInsets.all(10),
              child: ListTile(
                onTap: () {
                  Scaffold.of(context).showSnackBar(SnackBar(
                    content: Text(
                      "Выбрано блюдо '" + _dishList[index].name + '\'',
                      textAlign: TextAlign.center,
                      style: TextStyle(
                          fontSize: 16.0, fontWeight: FontWeight.bold),
                    ),
                    duration: Duration(seconds: 2),
                  ));
                },
                leading: Image.network("${_dishList[index].imageLink}${Random.secure().nextInt(100).toString()}"),
                title: Text(_dishList[index].name),
                subtitle: Text(_dishList[index].shortDescription),
                trailing: Icon(Icons.keyboard_arrow_right),
              ),
            );
          },
        ),
      ),
    );
  }

  List<DishItem> _loadDish(String categorySysName) {
    List<DishItem> tmp = [];

    for (var i = 1; i <= 100; i++) {
      tmp.add(new DishItem("имя " + i.toString(), "описание " + i.toString()));
    }
    return tmp;
  }
}
