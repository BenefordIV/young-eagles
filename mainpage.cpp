#include "mainpage.h"
#include "ui_mainpage.h"

#include <qmessagebox.h>

mainpage::mainpage(QWidget *parent)
    : QWidget(parent)
    , ui(new Ui::mainpage)
{
    ui->setupUi(this);
}

mainpage::~mainpage()
{
    delete ui;
}

void mainpage::openWindow(QSqlDatabase conn)
{
    this->show();
    ui->mainPage_Label->setText("Welcome " + conn.userName() + "!");
}
