#ifndef MAINPAGE_H
#define MAINPAGE_H

#include "databaseconn.h"

#include <QWidget>

namespace Ui {
class mainpage;
}

class mainpage : public QWidget
{
    Q_OBJECT

public:
    explicit mainpage(QWidget *parent = nullptr);
    ~mainpage();
    void openWindow(QSqlDatabase conn);

protected:


private:
    Ui::mainpage *ui;
};

#endif // MAINPAGE_H
