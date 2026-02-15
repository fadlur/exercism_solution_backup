#include <string>

namespace log_line {
    std::string message(std::string line) {
        // return the message
        int position = line.find(":") + 2;
        return line.substr(position, line.size() - position);
    }

    std::string log_level(std::string line) {
        // return the log level
        return line.substr(1, line.find(":") - 2);
    }

    std::string reformat(std::string line) {
        // return the reformatted message
        return message(line)+" ("+log_level(line)+")";
    }
}
