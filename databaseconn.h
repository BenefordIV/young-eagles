#ifndef DATABASECONN_H
#define DATABASECONN_H

#include <QObject>
#include <qsqldatabase.h>

class DatabaseConn : public QObject
{
    Q_OBJECT
public:
    explicit DatabaseConn(QObject *parent = nullptr);
    static QSqlDatabase CreateDatabaseConn(QString userName, QString password);

signals:
};

#endif // DATABASECONN_H
