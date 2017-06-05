import React, {PropTypes} from 'react';

const Sample = ({
    showData
}) => (
    <div>{showData}</div>
)

Sample.propTypes = {
    showData: PropTypes.string.isRequired
}

export default Sample
