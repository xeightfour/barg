#include <iostream>

#include <node.h>

using std::endl;
using std::cerr;
using std::clog;

using namespace barg;

Node::Node() : ID(0), type(kValve), status(kPassive), label("Node") {
    cerr << "[WARNING] -> Initializing empty node" << endl;
}

Node::Node(int iID, NodeType iType, Status iStatus, std::string_view iLabel) : ID(iID), type(iType), status(iStatus), label(iLabel) {
    clog << "[LOG] -> Initialized node " << iLabel << " with ID " << iID << endl;
}

void Node::addUpStream(int iNodeID, int iStreamID) {
    upStream.push_back(std::pair<int, int>(iNodeID, iStreamID));
    clog << "[LOG] -> Node " << ID << " routed stream " << iStreamID << " from node " << iNodeID << endl;
}

void Node::addDownStream(int iNodeID, int iStreamID) {
    downStream.push_back(std::pair<int, int>(iNodeID, iStreamID));
    clog << "[LOG] -> Node " << ID << " routed stream " << iStreamID << " to node " << iNodeID << endl;
}

void Node::setStatus(Status iStatus) {
    status = iStatus;
    clog << "[LOG] -> Set status " << iStatus << " for node " << ID << endl;
}

int Node::getID() const {
    return ID;
}

NodeType Node::getType() const {
    return type;
}

Status Node::getStatus() const {
    return status;
}

std::string_view Node::getLabel() const {
    return label;
}
