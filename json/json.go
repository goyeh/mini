package json

import (
    "bytes"
    "encoding/json"
)

const (
    empty = ""
    tab   = "\t"
)

// --------------------------------------------
//   _____ _        _             _  __
//  /  ___| |      (_)           (_)/ _|
//  \ `--.| |_ _ __ _ _ __   __ _ _| |_ _   _
//   `--. \ __| '__| | '_ \ / _` | |  _| | | |
//  /\__/ / |_| |  | | | | | (_| | | | | |_| |
//  \____/ \__|_|  |_|_| |_|\__, |_|_|  \__, |
//                           __/ |       __/ |
//                          |___/       |___/
// Returns the string representation of the JSON
func Stringify(data interface{}) (string, error) {
    b, err := json.Marshal(data)
    if err != nil {
        return empty, err
    }
    return string(b), nil
}



// ----------------------------------------------
//   _____ _                   _    _  __
//  /  ___| |                 | |  (_)/ _|
//  \ `--.| |_ _ __ _   _  ___| | ___| |_ _   _
//   `--. \ __| '__| | | |/ __| |/ / |  _| | | |
//  /\__/ / |_| |  | |_| | (__|   <| | | | |_| |
//  \____/ \__|_|   \__,_|\___|_|\_\_|_|  \__, |
//                                         __/ |
//                                        |___/
// Structify converts to the original structure
func Structify(data string, value interface{}) error {
    return json.Unmarshal([]byte(data), value)
}

// -----------------------------------
//     _                 _  __
//    (_)               (_)/ _|
//     _ ___  ___  _ __  _| |_ _   _
//    | / __|/ _ \| '_ \| |  _| | | |
//    | \__ \ (_) | | | | | | | |_| |
//    | |___/\___/|_| |_|_|_|  \__, |
//   _/ |                       __/ |
//  |__/                       |___/
//  jsonify returns a pretty json text
func Jsonify(data interface{}) (string, error) {
    buffer := new(bytes.Buffer)
    encoder := json.NewEncoder(buffer)
    encoder.SetIndent(empty, tab)

    err := encoder.Encode(data)
    if err != nil {
        return empty, err
    }
    return buffer.String(), nil
}