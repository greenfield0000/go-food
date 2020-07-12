import 'dart:math';
import 'dart:ui';

import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter/painting.dart';
import 'package:mobile/models/list/category/category-item.dart';

import 'dish-list-page.dart';

class MainPage extends StatefulWidget {
  MainPage({Key key, String tittle}) : super(key: key);

  @override
  _MainPageState createState() => _MainPageState();
}

class _MainPageState extends State<MainPage> {
  /*
  * Построение списка категорий
  */
  CustomScrollView buildCustomScrollView() {
    return CustomScrollView(
      primary: false,
      slivers: <Widget>[
        SliverPadding(
          padding: const EdgeInsets.all(20),
          sliver: SliverGrid.count(
            crossAxisSpacing: 10,
            mainAxisSpacing: 10,
            crossAxisCount: 2,
            children: <Widget>[
              _buildCategoryItem(
                  context, CategoryItem("Салаты", "системное имя 1")),
              _buildCategoryItem(
                  context, CategoryItem("Первые блюда", "системное имя 2")),
              _buildCategoryItem(
                  context, CategoryItem("Вторые блюда", "системное имя 3")),
              _buildCategoryItem(
                  context, CategoryItem("Гарниры", "системное имя 4")),
              _buildCategoryItem(
                  context, CategoryItem("Хлеб", "системное имя 5")),
              _buildCategoryItem(
                  context, CategoryItem("Выпечка", "системное имя 6")),
              _buildCategoryItem(context,
                  CategoryItem("Кондитерские изделия", "системное имя 7")),
              _buildCategoryItem(
                  context, CategoryItem("Торты", "системное имя 8"))
            ],
          ),
        ),
      ],
    );
  }

  /*
   * Создание категории меню
   */
  Builder _buildCategoryItem(BuildContext context, CategoryItem categoryItem) {
    String image = '${categoryItem.imageLink.toString()}${Random.secure().nextInt(100).toString()}';
    return Builder(
      builder: (context) => Container(
        child: InkWell(
          onTap: () {
            Navigator.push(
                context,
                MaterialPageRoute(
                    builder: (context) =>
                        DishListPage(categoryItem.name, categoryItem.sysName)));
          },
          child: Container(
            padding: const EdgeInsets.all(8),
            decoration: BoxDecoration(
              boxShadow: [
                BoxShadow(
                  color: Colors.grey.withOpacity(0.5),
                  spreadRadius: 5,
                  blurRadius: 7,
                  offset: Offset(0, 3), // changes position of shadow
                ),
              ],
              color: Colors.white,
              borderRadius: BorderRadius.all(Radius.circular(10.0)),
            ),
            child: Container(
              height: double.infinity,
              width: double.infinity,
              child: Column(
                children: [
                  Text(
                    categoryItem.name,
                    textAlign: TextAlign.center,
                    style: TextStyle(
                      fontWeight: FontWeight.bold
                    ),
                  ),
                  SizedBox(
                    height: 15,
                  ),
                  Image.network(
                    image,
                    fit: BoxFit.contain,
                  )
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          title: Text('Меню'),
        ),
        drawer: Drawer(
          child: ListView(
            padding: EdgeInsets.zero,
            children: [
              DrawerHeader(
                decoration: BoxDecoration(
                  color: Colors.blue,
                ),
                padding: EdgeInsets.all(10),
                child: CircleAvatar(
//                    radius: 10,
//                    backgroundColor: Color(0xffFDCF09),
                  child: CircleAvatar(
                    radius: 75,
                    backgroundImage: NetworkImage(
                        'https://sun9-57.userapi.com/c855628/v855628430/83a6e/7mr8s5PVxvY.jpg'),
                  ),
                ),
              ),
              _buildListTileDrawer(Icon(Icons.account_circle), Text('Профиль')),
              _buildListTileDrawer(Icon(Icons.message), Text('Сообщения')),
              _buildListTileDrawer(Icon(Icons.settings), Text('Настройки')),
            ],
          ),
        ),
        body: buildCustomScrollView());
  }

/*
 * Создание итема слайдер бара
 */
  ListTile _buildListTileDrawer(Icon icon, Text text) {
    return ListTile(
      leading: icon,
      title: text,
    );
  }
}
