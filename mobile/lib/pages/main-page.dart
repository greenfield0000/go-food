import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class MainPage extends StatefulWidget {
  MainPage({Key key, String tittle}) : super(key: key);

  @override
  _MainPageState createState() => _MainPageState();
}

class _MainPageState extends State<MainPage> {
  @override
  Widget build(BuildContext context) {
    return Container(
      child: Scaffold(
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
                  child: Text(
                    'Drawer',
                    style: TextStyle(
                      color: Colors.white,
                      fontSize: 24,
                    ),
                  ),
                ),
                _buildListTileDrawer(Icon(Icons.message), Text('Messages')),
                _buildListTileDrawer(
                    Icon(Icons.account_circle), Text('Profile')),
                _buildListTileDrawer(Icon(Icons.settings), Text('Settings')),
              ],
            ),
          ),
          body: buildCustomScrollView()),
    );
  }

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
              _buildCategoryItem("Салаты")
              ,_buildCategoryItem("Первые блюда")
              ,_buildCategoryItem("Вторые блюда")
              ,_buildCategoryItem("Гарниры")
              ,_buildCategoryItem("Хлеб")
              ,_buildCategoryItem("Выпечка")
              ,_buildCategoryItem("Кондитерские изделия")
              ,_buildCategoryItem("Торты")
            ],
          ),
        ),
      ],
    );
  }

  /*
   * Создание категории меню
   */
  Container _buildCategoryItem(String text) {
    return Container(
      padding: const EdgeInsets.all(8),
      child: Text(text.toString()),
      color: Colors.green[100],
    );
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
