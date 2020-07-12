/*
 * Аттрибут листа блюда
 */
class DishItem {
  // название блюда
  final String _name;

  // короткое описание блюда
  final String _shortDescription;

  // ссылка на изображение
  final String _imageLink = "https://picsum.photos/100?image=";

  DishItem(this._name, this._shortDescription);

  String get imageLink => _imageLink;

  String get shortDescription => _shortDescription;

  String get name => _name;
}
