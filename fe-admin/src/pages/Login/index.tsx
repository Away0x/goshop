import React, { FC, FormEvent } from 'react';
import {
  Form,
  Input,
  Button,
  message as messageAlert,
} from 'antd';
import { FormComponentProps } from 'antd/lib/form';

import './index.less';

import { withRouter, RouteComponentProps } from 'aw-react-router';
import { awRouter } from '@/routes';
import AuthStore from '@/store/auth';


const LoginForm: FC<FormComponentProps & RouteComponentProps> = ({ form, history }) => {
  const { loginAction } = AuthStore.useStore();

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();

    form.validateFields(async (err, values) => {
      if (err) { return; }
      const {status, message} = await loginAction(values['username'], values['password']);
      if (!status) {
        messageAlert.error(message || '登录失败');
        return;
      }

      const homeRoute = awRouter.find('home');
      history.push(homeRoute ? homeRoute.fullPath : '/');
    });
  };

  return (
    <Form className="login-form" layout="vertical" onSubmit={handleSubmit}>
      <h3 className="title">XXXXXX 平台</h3>
      <div className="form-body">
        <Form.Item className="form-item">
          {form.getFieldDecorator('username', {
            rules: [{ required: true, message: '请输入正确的用户名' }],
          })(
            <Input placeholder="Username" />,
          )}
        </Form.Item>
        <Form.Item className="form-item">
          {form.getFieldDecorator('password', {
            rules: [{ required: true, message: '请输入密码' }],
          })(
            <Input type="password" placeholder="Password" />,
          )}
        </Form.Item>
        <Form.Item className="form-item">
          <Button type="primary" htmlType="submit" className="form-button">
            登 录
        </Button>
        </Form.Item>
      </div>
    </Form>
  );
};

const WrappedLoginForm = Form.create({ name: 'login' })(withRouter(LoginForm));

const LoginPage: FC = () => {
  return (
    <div className="login-page">
      <WrappedLoginForm />
    </div>
  );
};

export default LoginPage;
