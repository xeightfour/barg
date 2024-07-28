#include <iostream>

#include <stream.h>
#include <node.h>
#include <graph.h>

using std::cin;
using std::endl;
using std::cout;

using namespace barg;

enum ExitCode {
    kExitFailure = -1,
    kExitSuccess
};

int main() {
    std::ios_base::sync_with_stdio(false);

    Graph G;

    G.setNodes({
        Node(0, kSource, kPassive, "Source"),
        Node(0, kValve, kPassive, "Valve"),
        Node(0, kSink, kPassive, "Sink")
    });

    G.setStreams({
        Stream(0, 0, 1, kActive),
        Stream(0, 1, 2, kPassive),
        Stream(0, 2, 0, kActive)
    });

    G.updateNodeStatus(0, kActive);
    G.updateNodeStatus(1, kActive);
    G.updateNodeStatus(2, kActive);

    bool status = G.validate();
    cout << status << endl;

    G.updateStreamStatus(1, kActive);
    status = G.validate();
    cout << status << endl;

    G.updateStreamStatus(2, kPassive);
    status = G.validate();
    cout << status << endl;

    return kExitSuccess;
}
