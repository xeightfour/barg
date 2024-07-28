#include <iostream>
#include <list>

#include <graph.h>

using std::endl;
using std::cerr;
using std::clog;

using namespace barg;

int Graph::count = 0;

Graph::Graph() : handle(Graph::count++), order(0), size(0) {
    clog << "[LOG] Initialized graph " << handle << endl;
}

int Graph::findSink(int node = 0) const {
    bool mark[order];
    std::fill(mark, mark + order, false);
    std::list chain { node };
    bool foundSink = false;
    while (!chain.empty()) {
        int onVertex = chain.front();
        chain.pop_front();
        mark[onVertex] = true;
        for (auto [toVertex, toStream] : nodes[onVertex].downStream) {
            if (streams[toStream].getStatus() == kPassive || nodes[toVertex].getStatus() == kPassive) {
                continue;
            }
            if (mark[toVertex]) {
                return -1;
            }
            foundSink |= (nodes[toVertex].getType() == kSink);
            chain.push_back(toVertex);
        }
    }
    return foundSink;
}

bool Graph::validate() const {
    bool res = true;
    for (int ID = 0; ID < size; ID++) {
        if (nodes[ID].getType() == kSource && nodes[ID].getStatus() == kActive) {
            int status = findSink(ID);
            if (status < 0) {
                clog << "[WARNING] Graph " << handle << " doesn't resemble a DAG" << endl;
                return false;
            } else if (status == 0) {
                clog << "[WARNING] Node " << ID << " of graph " << handle << " is a bomb!" << endl;
                res &= (bool)status;
            }
        }
    }
    return res;
}

int Graph::addNode(NodeType type, Status status, std::string_view label) {
    clog << "[LOG] Graph " << handle << " adding node " << order << endl;
    nodes.push_back(Node(order, type, status, label));
    return order++;
}

int Graph::addStream(int head, int tail, Status status) {
    if (head >= order || tail >= order) {
        cerr << "[ERROR] Stream nodes are out of bound" << endl;
    }
    clog << "[LOG] Graph " << handle << " adding stream " << size << endl;
    streams.push_back(Stream(size, head, tail, status));
    nodes[head].addDownStream(tail, size);
    nodes[tail].addUpStream(head, size);
    return size++;
}

void Graph::setNodes(std::vector<Node> iNodes) {
    for (Node node : iNodes) {
        addNode(node.getType(), node.getStatus(), node.getLabel());
    }
}

void Graph::setStreams(std::vector<Stream> iStreams) {
    for (Stream stream : iStreams) {
        addStream(stream.getHead(), stream.getTail(), stream.getStatus());
    }
}

void Graph::updateNodeStatus(int ID, Status status) {
    clog << "[LOG] Graph " << handle << " updating node " << ID << endl;
    if (ID >= order) {
        cerr << "[ERROR] No node with ID " << ID << " found" << endl;
        return;
    }
    nodes[ID].setStatus(status);
}

void Graph::updateStreamStatus(int ID, Status status) {
    clog << "[LOG] Graph " << handle << " updating stream " << ID << endl;
    if (ID >= size) {
        cerr << "[ERROR] No stream with ID " << ID << " found" << endl;
        return;
    }
    streams[ID].setStatus(status);
}

int Graph::getOrder() const {
    return order;
}

int Graph::getSize() const {
    return size;
}
