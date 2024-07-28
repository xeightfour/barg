#pragma once

#include <string_view>
#include <vector>

#include <barg.h>
#include <stream.h>
#include <node.h>

namespace barg {

    struct Graph {
        static int count;

        const int handle;
        int order;
        int size;

        std::vector<Node> nodes;
        std::vector<Stream> streams;

        Graph();

        int findSink(int node) const;
        bool validate() const;

        int addNode(NodeType type, Status status, std::string_view label);
        int addStream(int head, int tail, Status status);

        void setNodes(std::vector<Node> iNodes);
        void setStreams(std::vector<Stream> iStreams);

        void updateNodeStatus(int ID, Status status);
        void updateStreamStatus(int ID, Status status);

        int getOrder() const;
        int getSize() const;
    };

};
