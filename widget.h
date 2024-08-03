#ifndef WIDGET_H
#define WIDGET_H

#include "databaseconn.h"
#include <QWidget>

QT_BEGIN_NAMESPACE
namespace Ui {
class Widget;
}
QT_END_NAMESPACE

class Widget : public QWidget
{
    Q_OBJECT

public:
    Widget(QWidget *parent = nullptr);
    ~Widget();

private slots:
    void on_pushButton_Login_clicked();

private:
    Ui::Widget *ui;
    QSqlDatabase dbConn;
};
#endif // WIDGET_H
