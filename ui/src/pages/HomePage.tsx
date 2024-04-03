import React from "react";
import {Link} from "react-router-dom";

export default function HomePage(): React.ReactElement {
    return (<>
        <header>Home</header>
        <main>
            <p>Hello World</p>
            <Link to='/Login'>Login</Link>
        </main>
    </>);
}