import React from 'react';
import { LayoutProps } from 'aw-react-router';
import {
  Form,
  Input,
  Button,
  message as messageAlert,
} from 'antd';
import { FormComponentProps } from 'antd/lib/form';

import './index.less';

const LoginForm: React.FC<FormComponentProps> = ({ form }) => {

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    form.validateFields(async (err, values) => {
      if (err) { return; }
      // if (!status) {
      //   messageAlert.error(message || '登录失败');
      //   return;
      // }

    });
  };

  return (
    <Form className="login_form" layout="vertical" onSubmit={handleSubmit}>
      <h3 className="title">XXXXXX 平台</h3>
      <div className="form_body">
        <Form.Item className="form_item">
          {form.getFieldDecorator('username', {
            rules: [{ required: true, message: '请输入正确的用户名' }],
          })(
            <Input placeholder="Username" />,
          )}
        </Form.Item>
        <Form.Item className="form_item">
          {form.getFieldDecorator('password', {
            rules: [{ required: true, message: '请输入密码' }],
          })(
            <Input type="password" placeholder="Password" />,
          )}
        </Form.Item>
        <Form.Item className="form_item">
          <Button type="primary" htmlType="submit" className=".form_button">
            登 录
        </Button>
        </Form.Item>
      </div>
    </Form>
  );
};

const WrappedLoginForm = Form.create({ name: 'login' })(LoginForm);

const LoginPage: React.FC<LayoutProps> = () => {
  return (
    <div className="page_login">
      <WrappedLoginForm />
    </div>
  );
};

export default LoginPage;
