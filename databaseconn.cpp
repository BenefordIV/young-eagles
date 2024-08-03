#include "databaseconn.h"

#include <qsqldatabase.h>

DatabaseConn::DatabaseConn(QObject *parent)
    : QObject{parent}
{}

QSqlDatabase DatabaseConn::CreateDatabaseConn(QString userName, QString password)
{
    QSqlDatabase db = QSqlDatabase::addDatabase("QMYSQL");
    db.setHostName("localhost");
    db.setUserName(userName);
    db.setPassword(password);
    db.setPort(3306);
    db.setDatabaseName("young_eagles_dev");

    return db;
}
