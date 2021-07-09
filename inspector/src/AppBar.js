import * as React from "react";

import {
    Navbar,
    NavbarGroup,
    NavbarHeading,
} from "@blueprintjs/core";

function AppBar () {
    return (
        <Navbar>
            <NavbarGroup >
                <NavbarHeading>Inspector</NavbarHeading>
            </NavbarGroup>
        </Navbar>
    );
}

export default AppBar
