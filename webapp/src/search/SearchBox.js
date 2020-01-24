import React from "react";

export function SearchBox(props) {
    return <div className="row">
        <div className="input-field col s12">
            <input placeholder="Type skill name to filter..."
                   id="search_string" type="text" className="validate"
                   onChange={e => props.updateSearch(e.target.value)}
                   value={props.search}
            />
        </div>
    </div>
}