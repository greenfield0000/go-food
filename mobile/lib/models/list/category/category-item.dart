/*
 * Представление стуктуры элемента категории
 */
class CategoryItem {
  final String _name;
  final String _sysName;
  final String _imageLink = "https://picsum.photos/100?image=";

  CategoryItem(this._name, this._sysName);

  String get sysName => _sysName;

  String get name => _name;

  String get imageLink => _imageLink;
}
