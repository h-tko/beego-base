import {Component} from 'react';

export default class MainView extends Component {
    render() {
        return (
            <div className="row">
                <form>
                    <div className="col-xs-10">
                        <div className="input-group">
                            <span className="input-group-addon">名前</span>
                            <input type="text" className="form-control"></input>
                        </div>
                    </div>
                </form>
            </div>
        )
    }
}
