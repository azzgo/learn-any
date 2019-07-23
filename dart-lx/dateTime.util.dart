format(DateTime dateTime, String  format) {

  return format
    .replaceAll('YYYY', dateTime.year.toString())
    .replaceAll('MM', dateTime.month.toString())
    .replaceAll('DD', dateTime.day.toString())
    .replaceAll('HH', dateTime.hour.toString())
    .replaceAll('mm', dateTime.minute.toString())
    .replaceAll('ss', dateTime.second.toString());
}