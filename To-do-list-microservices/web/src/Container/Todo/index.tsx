import React from 'react';
import { Checkbox, Card } from 'antd';
import { CheckboxChangeEvent } from 'antd/lib/checkbox';

export const Todo = () => {
  const onChange = (e: CheckboxChangeEvent) => {
    console.log(`checked = ${e.target.checked}`);
  };
  //TODO check row display
  return (
    <div>
      <Card style={{ width: 400 }}>
        <Checkbox onChange={onChange}>Checkbox</Checkbox>
        <Checkbox onChange={onChange}>Checkbox</Checkbox>
        <Checkbox onChange={onChange}>Checkbox</Checkbox>
        <Checkbox onChange={onChange}>Checkbox</Checkbox>
      </Card>
    </div>
  );
};
