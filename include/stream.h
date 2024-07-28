#pragma once

#include <barg.h>

namespace barg {

    struct Stream {
        const int ID;
        const int head;
        const int tail;
        Status status;

        Stream();
        Stream(int iID, int iHead, int iTail, Status iStatus);

        void setStatus(Status iStatus);

        int getID() const;
        int getHead() const;
        int getTail() const;
        Status getStatus() const;
    };

};
