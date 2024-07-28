#include <iostream>

#include <stream.h>

using std::endl;
using std::cerr;
using std::clog;

using namespace barg;

Stream::Stream() : ID(0), head(0), tail(0), status(kPassive) {
    cerr << "[WARNING] -> Initializing empty stream" << endl;
}

Stream::Stream(int iID, int iHead, int iTail, Status iStatus) : ID(iID), head(iHead), tail(iTail), status(iStatus) {
    clog << "[LOG] -> Initialized stream " << iID << " from " << iHead << " to " << iTail << endl;
}

void Stream::setStatus(Status iStatus) {
    status = iStatus;
    clog << "[LOG] -> Set status " << iStatus << " for stream " << ID << endl;
}

int Stream::getID() const {
    return ID;
}

int Stream::getHead() const {
    return head;
}

int Stream::getTail() const {
    return tail;
}

Status Stream::getStatus() const {
    return status;
}
