#pragma once

namespace barg {

    enum NodeType {
        kSink = -1,
        kValve,
        kSource
    };

    enum Status {
        kPassive = 0,
        kActive
    };

};
