import 'dart:ui';

import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter/painting.dart';

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
              _buildCategoryItem(context, "Салаты"),
              _buildCategoryItem(context, "Первые блюда"),
              _buildCategoryItem(context, "Вторые блюда"),
              _buildCategoryItem(context, "Гарниры"),
              _buildCategoryItem(context, "Хлеб"),
              _buildCategoryItem(context, "Выпечка"),
              _buildCategoryItem(context, "Кондитерские изделия"),
              _buildCategoryItem(context, "Торты")
            ],
          ),
        ),
      ],
    );
  }

  /*
   * Создание категории меню
   */
  Builder _buildCategoryItem(BuildContext context, String text) {
    return Builder(
      builder: (context) => Container(
        child: InkWell(
          onTap: () {
            Scaffold.of(context).showSnackBar(SnackBar(
              content: Text(
                "Выбрана категория '" + text + '\'',
                textAlign: TextAlign.center,
                style: TextStyle(fontSize: 16.0, fontWeight: FontWeight.bold),
              ),
              duration: Duration(seconds: 2),
            ));
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
            child: Text(
              text.toString(),
              textAlign: TextAlign.center,
            ),
//            color: Colors.green[100],
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
