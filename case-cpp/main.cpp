#include <stdio.h>
#include <string>
#include <iostream>
#include <fstream>
#include "window.pb.h"

using namespace std;
using namespace protobuf;

int write(const string &filename)
{
    lnks::WindowId winId1;
    lnks::MainWindowList winList;
    winId1.chromium_id(1);
    winList.ok(true);
    winList.add_windows(winId1);

    printf("write to %s\n", filename.c_str());

    // Write the data to disk.
    fstream output(filename, ios::out | ios::trunc | ios::binary);

    if (!winList.SerializeToOstream(&output)) {
        cerr << "Failed to write." << endl;
        return -1;
    }
}

int read(const string &filename)
{
    lnks::MainWindowList winList;

    printf("read from %s\n", filename.c_str());
    {
        fstream input(filename, ios::in | ios::binary);
        if (!winList.ParseFromIstream(&input)) {
            cerr << "Failed to parse." << endl;
            return -1;
        }
    }

    cout << "id: " << winList.ok() << endl;
    cout << "str: " << winList.windows().chromium_id() << endl;

    return 0;
}

int main(int argc, const char *argv[])
{
    write("test.data");
    read("test.data");

	return 0;
}
