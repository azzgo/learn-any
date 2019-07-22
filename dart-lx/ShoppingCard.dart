import 'Item.dart';

class ShoppingCard {
  String name;
  DateTime dateTime;
  String code;

  List<Item> bookings;

  ShoppingCard(this.name, this.code): dateTime = DateTime.now();

  price() {
    double sum = 0.0;

    for (var i in bookings) {
      sum + i.price;
    }

    return sum;
  }

  getInfo() {
    return '''购物车信息:
------------------------
用户名: $name
优惠码: $code
总价:   ${price().toString()}
日期:   ${dateTime.toString()}
------------------------
    ''';
  }
}
