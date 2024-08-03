#include "widget.h"
#include "./ui_widget.h"
#include <QSqlDatabase>
#include <QMessageBox>

Widget::Widget(QWidget *parent)
    : QWidget(parent)
    , ui(new Ui::Widget)
{
    ui->setupUi(this);
}

Widget::~Widget()
{
    delete ui;
}

void Widget::on_pushButton_Login_clicked()
{
    dbConn = DatabaseConn::CreateDatabaseConn(ui->lineEdit_Username->text(), ui->lineEdit_Password->text());

    QMessageBox mb;

    if(!dbConn.open()) {
        mb.setText("Unable to connect to db");
        mb.exec();
    } else {
        mb.setText("connected to db");
        mb.exec();
    }
}

