import {connect} from 'react-redux';
import React, {Component, PropTypes} from 'react';
import {bindActionCreators} from 'redux';

import Sample from '../components/Sample';
import * as sampleActions from '../actions/sample';

class SampleContainer extends Component {
    constructor(props) {
        super(props);
    }

    render() {
        const {showData} = this.props;

        return (
            <div>
                <Sample showData={showData}></Sample>
            </div>
        );
    }
};

function mapStateToProps(state) {
    const {showData} = this.props.sample;
    return {
        showData,
    };
}

function mapDispatchToProps(dispatch) {
    return {
        sampleActionBind: bindActionCreators(sampleActions, dispatch)
    };
}

export default connect (
    mapStateToProps,
    mapDispatchToProps
)(SampleContainer);
