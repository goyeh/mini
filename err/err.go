package err

import "errors"

// ------------------------------------
//   _   _
//  | \ | |
//  |  \| | _____      __
//  | . ` |/ _ \ \ /\ / /
//  | |\  |  __/\ V  V /
//  \_| \_/\___| \_/\_/
// New returns an error with the text.
func New(text string) (New error) {
    New = errors.New(text)
    return
}

// --------------------------------------------
//   _____                  _____ __ _   _ _ _
//  |  ___|                |_   _/ _| \ | (_) |
//  | |__ _ __ _ __ ___  _ __| || |_|  \| |_| |
//  |  __| '__| '__/ _ \| '__| ||  _| . ` | | |
//  | |__| |  | | | (_) | | _| || | | |\  | | |
//  \____/_|  |_|  \___/|_| \___/_| \_| \_/_|_|
// ErrorIfNil checks if the err is nil, if true
// returns the passed message otherwise err.Error()
func ErrorIfNil(err error, defaultMessage string) (ErrorIfNil string) {
    if err != nil {
        return err.Error()
    }
    ErrorIfNil = defaultMessage
    return
}
