#pragma once

#include <string_view>
#include <utility>
#include <vector>

#include <barg.h>

namespace barg {

    struct Node {
        const int ID;
        const NodeType type;
        Status status;
        std::string_view label;

        std::vector<std::pair<int, int>> upStream;
        std::vector<std::pair<int, int>> downStream;

        Node();
        Node(int iID, NodeType iType, Status iStatus, std::string_view iLabel);

        void addUpStream(int iNodeID, int iStreamID);
        void addDownStream(int iNodeID, int iStreamID);

        void setStatus(Status iStatus);

        int getID() const;
        NodeType getType() const;
        Status getStatus() const;
        std::string_view getLabel() const;
    };

};
